using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GUI.Models;

public class UserInfo
{
    public string auth_token { get; set; } = string.Empty;
    public string name { get; set; } = string.Empty;
    public string link { get; set; } = string.Empty;
    public string about { get; set; } = string.Empty;
    public string last_online { get; set; } = string.Empty;
    public string image_id { get; set; } = string.Empty;
    public string id { get; set;} = string.Empty;

    public UserInfo() { }

    public UserInfo(string auth_token, string name, string link, string about, string last_online, string image_id, string id = "") 
    { 
        this.auth_token = auth_token;
        this.name = name;
        this.link = link;
        this.about = about;
        this.last_online = last_online;
        this.image_id = image_id;
        this.id = id;
    }

    public override string ToString()
    {
        return JsonConvert.SerializeObject(this);
    }

}
