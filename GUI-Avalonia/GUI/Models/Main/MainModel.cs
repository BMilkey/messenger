using ReactiveUI;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace GUI.Models.Main;

public class MainModel : ReactiveObject
{
    



    private ObservableCollection<UserInfo> users = new()
    {
        new UserInfo("auth_1", "name_1", "link_1", "about_1", "last_1", "image_1"),
        new UserInfo("auth_2", "name_2", "link_2", "about_2", "last_2", "image_2"),
        new UserInfo("auth_3", "name_3", "link_3", "about_3", "last_3", "image_3"),
        new UserInfo("auth_4", "name_4", "link_4", "about_4", "last_4", "image_4"),
        new UserInfo("auth_1", "name_1", "link_1", "about_1", "last_1", "image_1"),
        new UserInfo("auth_2", "name_2", "link_2", "about_2", "last_2", "image_2"),
        new UserInfo("auth_3", "name_3", "link_3", "about_3", "last_3", "image_3"),
        new UserInfo("auth_4", "name_4", "link_4", "about_4", "last_4", "image_4"),
        new UserInfo("auth_1", "name_1", "link_1", "about_1", "last_1", "image_1"),
        new UserInfo("auth_2", "name_2", "link_2", "about_2", "last_2", "image_2"),
        new UserInfo("auth_3", "name_3", "link_3", "about_3", "last_3", "image_3"),
        new UserInfo("auth_4", "name_4", "link_4", "about_4", "last_4", "image_4"),
        new UserInfo("auth_1", "name_1", "link_1", "about_1", "last_1", "image_1"),
        new UserInfo("auth_2", "name_2", "link_2", "about_2", "last_2", "image_2"),
        new UserInfo("auth_3", "name_3", "link_3", "about_3", "last_3", "image_3"),
        new UserInfo("auth_4", "name_4", "link_4", "about_4", "last_4", "image_4"),
        new UserInfo("auth_1", "name_1", "link_1", "about_1", "last_1", "image_1"),
        new UserInfo("auth_2", "name_2", "link_2", "about_2", "last_2", "image_2"),
        new UserInfo("auth_3", "name_3", "link_3", "about_3", "last_3", "image_3"),
        new UserInfo("auth_4", "name_4", "link_4", "about_4", "last_4", "image_4"),
    };

    public ObservableCollection<UserInfo> Users
    {
        get => users;
        set => this.RaiseAndSetIfChanged(ref users, value);
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
    }


}
