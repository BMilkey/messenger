using Newtonsoft.Json;
using ReactiveUI;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Diagnostics;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Threading;
using System.Threading.Tasks;

namespace GUI.Models.Main;

public class MainModel : ReactiveObject, IDisposable
{

    private CancellationTokenSource UsersByNameCts = new();
    private ObservableCollection<UserInfo> users = new();
    public ObservableCollection<UserInfo> Users
    {
        get => users;
        set => this.RaiseAndSetIfChanged(ref users, value);
    }

    private ObservableCollection<ChatInfo> chats = new();
    public ObservableCollection<ChatInfo> Chats
    {
        get => chats;
        set => this.RaiseAndSetIfChanged(ref chats, value);
    }

    public readonly string URL;
    public UserInfo UserInfo { get; private set; }
    private HttpClient client;

    private string urlParameters;

    private ChatInfo currentChat = new("","","","","","Choose any chat","");
    public ChatInfo CurrentChat
    {
        get => currentChat;
        set => this.RaiseAndSetIfChanged(ref currentChat, value);
    }

    private WebSocketListener wsListener;

    public void Dispose()
    {
        client.Dispose();
        wsListener.Dispose();
    }

    public MainModel(UserInfo userInfo, string URL, string urlParameters = "")
    {
        this.UserInfo = userInfo;
        this.URL = URL;
        this.urlParameters = urlParameters;
        client = new HttpClient();
        wsListener = new WebSocketListener(URL, userInfo.auth_token);

        client.BaseAddress = new Uri($"http://{URL}/");
        _ = FindUsersByName("");
        _ = FetchChats();
        _ = ReadReceivedMessages();
        _ = ReadReceivedChats();
    }


    private async Task ReadReceivedMessages()
    {
        while (true)
        {
            try
            {
                var receivedMsg = await wsListener.MessagesChannel.Reader.ReadAsync();
                Debug.WriteLine($"MainModel received msg: {receivedMsg}");
                var message = receivedMsg ?? new MessageInfo();
                Chats.First(x => x.id == message.chat_id).AddMessage(message);
            }
            catch (Exception ex)
            {
                Debug.WriteLine(ex);
                return;
            }
        }
    }

    private async Task ReadReceivedChats()
    {
        while (true)
        {
            try
            {
                var receivedChat = await wsListener.ChatsChannel.Reader.ReadAsync();
                Debug.WriteLine($"MainModel received chat: {receivedChat}");
                var chat = receivedChat ?? new ChatInfo();
                Chats.Add(chat);
            }
            catch (Exception ex)
            {
                Debug.WriteLine(ex);
                return;
            }
        }
    }


    private class UserByNameRequest
    {
        public string auth_token { get; set; } = string.Empty;
        public string name { get; set; } = string.Empty;

        public UserByNameRequest(string auth_token, string name)
        {
            this.auth_token = auth_token;
            this.name = name;
        }
    }

    private class UsersResponse
    {
        public List<UserInfo> Users { get; set; } = new();
    }

    public async Task FindUsersByName(string name)
    {
        UsersByNameCts.Cancel();
        UsersByNameCts = new();
        var ct = UsersByNameCts.Token;
        var request = new UserByNameRequest(UserInfo.auth_token, name);

        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/users_by_name", request);
        response.EnsureSuccessStatusCode();

        var usersResponse = await response.Content.ReadFromJsonAsync<UsersResponse>(cancellationToken: ct);

        if (ct.IsCancellationRequested)
            return;


        Users = new(usersResponse.Users ?? new List<UserInfo>());
    }

    private class CreateChatRequest
    {
        public string auth_token { get; set; } = string.Empty;
        public string title { get; set; } = string.Empty;
        public List<string> users_links { get; set; } = new();
        public CreateChatRequest(string auth_token, string title, IEnumerable<string> users_links)
        {
            this.auth_token = auth_token;
            this.title = title;
            this.users_links = new(users_links);
        }
    }

    private class CreateChatResponse
    {
        public string chat_id { get; set; } = string.Empty;
        public List<UserInfo> users { get; set; } = new();
    }

    public async Task CreateChat(string title, IEnumerable<UserInfo> users)
    {

        var request = new CreateChatRequest(UserInfo.auth_token, title, users.Select(x => x.link));

        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/create_chat_return_users", request);
        response.EnsureSuccessStatusCode();

        var createChatResponse = await response.Content.ReadFromJsonAsync<CreateChatResponse>();

        //Debug.WriteLine(JsonConvert.SerializeObject(createChatResponse));

        //_ = FetchChats();
    }

    private class GetChatsByAuthTokenRequest
    {
        public string auth_token { get; set; }
        public GetChatsByAuthTokenRequest(string auth_token)
        {
            this.auth_token = auth_token;
        }
    }

    private class GetChatsByAuthTokenResponse
    {
        public List<ChatInfo> chats { get; set; } = new();
    }

    public async Task FetchChats()
    {
        var request = new GetChatsByAuthTokenRequest(UserInfo.auth_token);
        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/chats_by_token", request);
        response.EnsureSuccessStatusCode();

        var getChatsResponse = await response.Content.ReadFromJsonAsync<GetChatsByAuthTokenResponse>();

        //Debug.WriteLine(JsonConvert.SerializeObject(getChatsResponse));

        Chats = new(getChatsResponse.chats ?? new List<ChatInfo>());
        foreach (var chat in Chats)
        {
            _ = FetchMessagesByChatId(chat.id);
        }
    }


    private class GetMessagesByChatIdRequest
    {
        public string auth_token;
        public string chat_id;
        public string from_date;
        public string to_date;
        public GetMessagesByChatIdRequest(string auth_token, string chat_id, string from_date, string to_date)
        {
            this.auth_token = auth_token;
            this.chat_id = chat_id;
            this.from_date = from_date;
            this.to_date = to_date;
        }
    }

    private class GetMessagesByChatIdResponse
    {
        public List<MessageInfo> messages { get; set; } = new();
    }

    public async Task FetchMessagesByChatId(string chatId)
    {
        await FetchMessagesByChatId(chatId, DateTime.MinValue, DateTime.MaxValue);
    }

    public async Task FetchMessagesByChatId(string chatId, DateTime From, DateTime To)
    {
        var request = new GetMessagesByChatIdRequest(UserInfo.auth_token, chatId, ToTimestamp(From), ToTimestamp(To));
        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/messages_by_chat_id", request);


        response.EnsureSuccessStatusCode();


        var getMessagesByChatIdResponse = await response.Content.ReadFromJsonAsync<GetMessagesByChatIdResponse>();

        var messages = new List<MessageInfo>(getMessagesByChatIdResponse.messages ?? new List<MessageInfo>());

        Chats.First(x => x.id == chatId).UpdateMessages(messages);

        //Debug.WriteLine(JsonConvert.SerializeObject(Chats.First(x => x.id == chatId)));

    }

    private static string ToTimestamp(DateTime dateTime)
    {
        dateTime = dateTime.ToUniversalTime();
        return dateTime.ToString("O");
    }

    private class CreateMessageRequest
    {
        public string auth_token = string.Empty;
        public string chat_id = string.Empty;
        public string reply_msg_id = string.Empty;
        public string text = string.Empty;
        public CreateMessageRequest(string auth_token, string chat_id, string reply_msg_id, string text)
        {
            this.auth_token = auth_token;
            this.chat_id = chat_id;
            this.reply_msg_id = reply_msg_id;
            this.text = text;
        }
    }

    private class CreateMessageResponse
    {
        public MessageInfo message { get; set; }
        public MessageInfo reply_msg { get; set; }
    }

    public async Task CreateMessage(string text, string reply_msg_id = "fake")
    {
        var request = new CreateMessageRequest(UserInfo.auth_token, CurrentChat.id, reply_msg_id, text);
        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/create_message", request);


        response.EnsureSuccessStatusCode();


        var CreateMessageResponse = await response.Content.ReadFromJsonAsync<CreateMessageResponse>();
        //Debug.WriteLine(JsonConvert.SerializeObject(CreateMessageResponse));
    }

}
