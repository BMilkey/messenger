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

public class MainModel : ReactiveObject
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


    public MainModel(UserInfo userInfo, string URL, string urlParameters = "")
    {
        this.UserInfo = userInfo;
        this.URL = URL;
        this.urlParameters = urlParameters;
        client = new HttpClient();
        client.BaseAddress = new Uri($"http://{URL}/");
        _ = FindUsersByName("");
        _ = GetChats(UserInfo.auth_token);
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

        Debug.WriteLine(JsonConvert.SerializeObject(createChatResponse));
        // TODO
        // add to chats
        // get_messages for this chat
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

    public async Task GetChats(string auth_token)
    {
        var request = new GetChatsByAuthTokenRequest(auth_token);
        client.DefaultRequestHeaders.Accept.Add(
        new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/chat/chats_by_token", request);
        response.EnsureSuccessStatusCode();

        var getChatsResponse = await response.Content.ReadFromJsonAsync<GetChatsByAuthTokenResponse>();

        Debug.WriteLine(JsonConvert.SerializeObject(getChatsResponse));
        // TODO
        // update chats
        Chats = new(getChatsResponse.chats ?? new List<ChatInfo>());
        // get_messages for every chat
    }
    // TODO
    //public async Task GetMessegesForChat...




}
