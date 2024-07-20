namespace GUI.Models;

public class MessageInfo
{
    public string chat_id { get; set; } = string.Empty;
    public string create_time { get; set; } = string.Empty;
    public string id { get; set; } = string.Empty;
    public string reply_msg_id { get; set; } = string.Empty;
    public string text { get; set; } = string.Empty;
    public string user_id { get; set; } = string.Empty;

    public MessageInfo(string chat_id, string create_time, string id, string reply_msg_id, string text, string user_id)
    {
        this.chat_id = chat_id;
        this.create_time = create_time;
        this.id = id;
        this.reply_msg_id = reply_msg_id;
        this.text = text;
        this.user_id = user_id;
    }
}
