using GUI.Models;
using Newtonsoft.Json;
using ReactiveUI;
using System.Collections.ObjectModel;
using System;
using System.Linq;
using AsyncImageLoader;
using GUI.Models.Main;
using AsyncImageLoader.Loaders;
using Newtonsoft.Json.Linq;
using System.Threading.Tasks;
using GUI.ViewModels.Main;

namespace GUI.ViewModels;

public class MainViewModel : ViewModelBase
{
    #region xaml fields
    public string Greeting => $"Welcome to Messenger, {mainModel.UserInfo.name}! Your's auth token is \"{mainModel.UserInfo.auth_token}\".";


    private string logsOutputField = "Main view logs will be here...\n";
    public string LogsOutputField
    {
        get => logsOutputField;
        set => this.RaiseAndSetIfChanged(ref logsOutputField, logsOutputField + "\n" + value);//logsOutputField + "\n" + value);
    }


    #endregion

    private readonly MainModel mainModel;
    public MainModel MainModel => mainModel;

    private ChatsViewModel chatsViewModel;
    public ChatsViewModel ChatsViewModel
    {
        get => chatsViewModel;
        set => this.RaiseAndSetIfChanged(ref chatsViewModel, value);
    }

    public MainViewModel(UserInfo userInfo, string URL) 
    {
        mainModel = new MainModel(userInfo, URL);
        LogsOutputField = $"Succesfully initialized with UserInfo: {userInfo.ToString()}";
        
        foreach (var item in mainModel.Users)
        {
            LogsOutputField = item.ToString();
            mainModel = new(userInfo, URL);
        }
        chatsViewModel = new ChatsViewModel(mainModel);

    }

    public MainViewModel()
    {
        LogsOutputField = "Incorrect initialization, UserInfo is empty";
        


    }
}
