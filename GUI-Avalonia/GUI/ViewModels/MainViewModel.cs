using GUI.Models;
using Newtonsoft.Json;
using ReactiveUI;
using System.Collections.ObjectModel;
using System;
using System.Linq;
using AsyncImageLoader;
using GUI.Models.Main;
using AsyncImageLoader.Loaders;

namespace GUI.ViewModels;

public class MainViewModel : ViewModelBase
{
    #region xaml fields
    public string Greeting => $"Welcome to Messenger, {mainModel.UserInfo.name}! Your's auth token is \"{mainModel.UserInfo.auth_token}\".";

    private readonly MainModel mainModel;
    public MainModel MainModel => mainModel;


    private string logsOutputField = "Main view logs will be here...\n";
    public string LogsOutputField
    {
        get => logsOutputField;
        set => this.RaiseAndSetIfChanged(ref logsOutputField, logsOutputField + "\n" + value);//logsOutputField + "\n" + value);
    }
    #endregion

    private IAsyncImageLoader imageLoader = new RamCachedWebImageLoader();
    public IAsyncImageLoader ImageLoader => imageLoader;

    public MainViewModel(UserInfo userInfo, string URL) 
    {
        mainModel = new MainModel(userInfo, URL);
        LogsOutputField = $"Succesfully initialized with UserInfo: {userInfo.ToString()}";
        foreach (var item in mainModel.Users)
        {
            LogsOutputField = item.ToString();
            mainModel = new(userInfo, URL);
        }
    }

    public MainViewModel()
    {
        LogsOutputField = "Incorrect initialization, UserInfo is empty";
        


    }
}
