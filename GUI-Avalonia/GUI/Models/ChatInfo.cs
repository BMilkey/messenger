using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GUI.Models;

public class ChatInfo
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

    public override string ToString()
    {
        return JsonConvert.SerializeObject(this);
    }
}
