using Newtonsoft.Json;
using ReactiveUI;
using System.Collections.Generic;
using System.Collections.ObjectModel;

namespace GUI.Models;

public class ChatInfo : ReactiveObject
{

    /*     "about": "string",
          "create_time": "string",
          "id": "string",
          "image_id": "string",
          "link": "string",
          "title": "string",
          "user_id": "string"
    */
    public string about { get; set; } = string.Empty;
    public string create_time { get; set; } = string.Empty;
    public string id { get; set; } = string.Empty;
    public string image_id { get; set; } = string.Empty;
    public string link { get; set; } = string.Empty;
    public string title { get; set; } = string.Empty;
    public string user_id { get; set; } = string.Empty;
    private ObservableCollection<MessageInfo> _messages = new ObservableCollection<MessageInfo>();
    public ObservableCollection<MessageInfo> messages
    {
        get => _messages;
        set => this.RaiseAndSetIfChanged(ref _messages, value);
    }
    public ChatInfo(string about, string create_time, string id, string image_id, string link, string title, string user_id)
    {
        this.about = about;
        this.create_time = create_time;
        this.id = id;
        this.image_id = image_id;
        this.link = link;
        this.title = title;
        this.user_id = user_id;
    }

    public void AddMessage(MessageInfo message)
    {
        messages.Add(message);
    }


    public void UpdateMessages(IEnumerable<MessageInfo> overwriteMessages)
    {
        messages = new(overwriteMessages);
    }

    public override string ToString()
    {
        return JsonConvert.SerializeObject(this);
    }
}
