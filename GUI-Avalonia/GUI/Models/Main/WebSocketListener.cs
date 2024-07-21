using Newtonsoft.Json;
using System;
using System.Diagnostics;
using System.Threading;
using System.Threading.Channels;
using System.Threading.Tasks;
using Websocket.Client;

namespace GUI.Models.Main;

public class WebSocketListener : IDisposable
{
    private WebsocketClient wsMessagesSubscriber;
    private readonly string URL;

    public Channel<MessageInfo> messagesChannel { get; private set; } = Channel.CreateUnbounded<MessageInfo>();
    public WebSocketListener(string URL, string auth_token)
    {
        this.URL = URL;
        wsMessagesSubscriber = new(new Uri($"ws://{URL}/sockets/subscribe_message_created"));
        _ = StartMessagesReceive(auth_token);
    }

    public void Dispose()
    {
        wsMessagesSubscriber.Dispose();
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

        wsMessagesSubscriber.MessageReceived.Subscribe(async respMsg => await HandleReceivedMessage(respMsg));
        wsMessagesSubscriber.DisconnectionHappened.Subscribe(async _ =>
        {
            //await wsMessagesSubscriber.Stop(System.Net.WebSockets.WebSocketCloseStatus.NormalClosure, "");
            //await wsMessagesSubscriber.NativeClient.CloseAsync(System.Net.WebSockets.WebSocketCloseStatus.NormalClosure, null, default);
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

            var receivedMsg = JsonConvert.DeserializeObject<MessageInfo>(respMsg.Text);
            Debug.WriteLine($"Received message: {respMsg.Text}");

            if (receivedMsg == null)
            {
                Debug.WriteLine("Received message is null");
            }
            await messagesChannel.Writer.WriteAsync(receivedMsg).ConfigureAwait(false);
        }
        catch (Exception ex)
        {
            Debug.WriteLine($"{ex}");
        }

    }
}
