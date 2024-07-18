using GUI.Models;
using Newtonsoft.Json;
using ReactiveUI;

namespace GUI.ViewModels;

public class MainViewModel : ViewModelBase
{
    #region xaml fields
    public string Greeting => $"Welcome to Messenger, {UserInfo.name}! Your's auth token is \"{UserInfo.auth_token}\".";
    private UserInfo userInfo;
    public UserInfo UserInfo
    {
        get => userInfo;
        set => this.RaiseAndSetIfChanged(ref userInfo, value);
    }

    private string logsOutputField = "Main view logs will be here...\n";
    public string LogsOutputField
    {
        get => logsOutputField;
        set => this.RaiseAndSetIfChanged(ref logsOutputField, logsOutputField + "\n" + value);//logsOutputField + "\n" + value);
    }
    #endregion

    public MainViewModel(UserInfo userInfo) 
    {
        this.userInfo = userInfo;
        LogsOutputField = $"Succesfully initialized with UserInfo: {JsonConvert.SerializeObject(userInfo)}";
    }

    public MainViewModel()
    {
        userInfo = new UserInfo();
        LogsOutputField = "Incorrect initialization, UserInfo is empty";
    }
}
