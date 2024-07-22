using Newtonsoft.Json;
using System;
using System.Diagnostics;
using System.Threading.Channels;
using System.Threading.Tasks;
using Websocket.Client;

namespace GUI.Models.Main;

public class WebSocketListener : IDisposable
{
    private WebsocketClient wsMessagesSubscriber;
    private WebsocketClient wsChatsSubscriber;
    private readonly string URL;

    public Channel<MessageInfo> MessagesChannel { get; private set; } = Channel.CreateUnbounded<MessageInfo>();
    public Channel<ChatInfo> ChatsChannel { get; private set; } = Channel.CreateUnbounded<ChatInfo>();
    public WebSocketListener(string URL, string auth_token)
    {
        this.URL = URL;
        wsMessagesSubscriber = new(new Uri($"ws://{URL}/sockets/subscribe_message_created"));
        wsChatsSubscriber = new(new Uri($"ws://{URL}/sockets/subscribe_сhat_сreated"));
        _ = StartMessagesReceive(auth_token);
        _ = StartChatsReceive(auth_token);
    }

    public void Dispose()
    {
        wsMessagesSubscriber.Dispose();
        wsChatsSubscriber.Dispose();
    }

    private async Task StartChatsReceive(string auth_token)
    {
        wsChatsSubscriber.ReconnectTimeout = null;// TimeSpan.FromSeconds(15);
        wsChatsSubscriber.ReconnectionHappened.Subscribe(info =>
        {
            wsChatsSubscriber.Send(auth_token);
            Debug.WriteLine($"Reconnection happened, type: {info.Type}");
        }
            );

        wsChatsSubscriber.MessageReceived.Subscribe(async respMsg =>
        {
            await HandleReceivedChat(respMsg);

        });
        wsChatsSubscriber.DisconnectionHappened.Subscribe(async _ =>
        {
            Debug.WriteLine($"{nameof(wsChatsSubscriber)} disconnected");
        });

        _ = wsChatsSubscriber.Start();

        _ = Task.Run(() => wsChatsSubscriber.Send(auth_token));
    }

    private class ChatInfoWs
    {
        public string id { get; set; }
        public string link { get; set; }
        public string title { get; set; }
        public string user_id { get; set; }
        public string create_time { get; set; }
        public string about { get; set; }
        public string image_id { get; set; }
    }

    private async Task HandleReceivedChat(ResponseMessage respMsg)
    {
        try
        {
            Debug.WriteLine($"Type of message: {respMsg.MessageType}");
            if (respMsg.MessageType != System.Net.WebSockets.WebSocketMessageType.Text)
                return;

            var parsedChat = JsonConvert.DeserializeObject<ChatInfoWs>(respMsg.Text);
            Debug.WriteLine($"Received chat: {respMsg.Text}");

            if (parsedChat == null)
            {
                Debug.WriteLine("Received chat is null");
            }
            var viewChat = new ChatInfo(parsedChat.about, parsedChat.create_time, parsedChat.id, parsedChat.image_id, parsedChat.link, parsedChat.title, parsedChat.user_id);
            await ChatsChannel.Writer.WriteAsync(viewChat).ConfigureAwait(false);
            wsChatsSubscriber.Send($"Got chat with id {parsedChat.id}");
        }
        catch (Exception ex)
        {
            Debug.WriteLine($"{ex}");
        }

    }

    private async Task StartMessagesReceive(string auth_token)
    {
        wsMessagesSubscriber.ReconnectTimeout = null;// TimeSpan.FromSeconds(15);
        wsMessagesSubscriber.ReconnectionHappened.Subscribe(info =>
            {
                wsMessagesSubscriber.Send(auth_token);
                Debug.WriteLine($"Reconnection happened, type: {info.Type}");
            }
            );

        wsMessagesSubscriber.MessageReceived.Subscribe(async respMsg =>
        {
            await HandleReceivedMessage(respMsg);
            
        });
        wsMessagesSubscriber.DisconnectionHappened.Subscribe(async _ =>
        {
            Debug.WriteLine($"{nameof(wsMessagesSubscriber)} disconnected");
        });

        _ = wsMessagesSubscriber.Start();

        _ = Task.Run(() => wsMessagesSubscriber.Send(auth_token));

    }

    private async Task HandleReceivedMessage(ResponseMessage respMsg)
    {

        try
        {
            Debug.WriteLine($"Type of msg: {respMsg.MessageType}");
            if (respMsg.MessageType != System.Net.WebSockets.WebSocketMessageType.Text)
                return;

            var parsedMsg = JsonConvert.DeserializeObject<MessageInfo>(respMsg.Text);
            Debug.WriteLine($"Received message: {respMsg.Text}");

            if (parsedMsg == null)
            {
                Debug.WriteLine("Received message is null");
            }
            await MessagesChannel.Writer.WriteAsync(parsedMsg).ConfigureAwait(false);
            wsMessagesSubscriber.Send($"Got message with id {parsedMsg.id}");
        }
        catch (Exception ex)
        {
            Debug.WriteLine($"{ex}");
        }

    }
}
