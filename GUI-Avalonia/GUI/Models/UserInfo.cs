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

}
