using Avalonia.Controls;
using GUI.Models;
using ReactiveUI;
using System.Diagnostics;
using System.IO;
using System.Net;
using System;
using System.Threading;
using System.Threading.Tasks;
using System.Windows.Input;
using System.Text.Json;

namespace GUI.ViewModels;

public class LoginViewModel : ViewModelBase
{
    #region xaml fields


    private string serverField = "147.45.70.245:8088";
    //"127.0.0.1:8081";//"147.45.70.245:8088";
    public string ServerField 
    {
        get => serverField;
        set => this.RaiseAndSetIfChanged(ref serverField, value);
    }

    private string loginField = string.Empty;
    public string LoginField 
    {
        get => loginField;
        set => this.RaiseAndSetIfChanged(ref loginField, value);
    }
    private string passwordField = string.Empty;
    public string PasswordField
    {
        get => passwordField;
        set => this.RaiseAndSetIfChanged(ref passwordField, value);
    }
    private string nameField = string.Empty;
    public string NameField
    {
        get => nameField;
        set => this.RaiseAndSetIfChanged(ref nameField, value);
    }
    private bool isAuth = true;
    public bool IsAuth
    {
        get => isAuth;
        set => this.RaiseAndSetIfChanged(ref isAuth, value);
    }

    private bool isAuthInProcess = true;
    public bool IsAuthInProcess
    {
        get => isAuthInProcess;
        set => this.RaiseAndSetIfChanged(ref isAuthInProcess, value);
    }

    public ICommand SendAuthCommand { get; set; }
    public ICommand ServerPingCommand { get; set; }

    private string logsOutputField = "Auth view logs will be here...\n";
    public string LogsOutputField
    {
        get => logsOutputField;
        set => this.RaiseAndSetIfChanged(ref logsOutputField, logsOutputField + "\n" + value);//logsOutputField + "\n" + value);
    }

    #endregion

    private UserInfo user = new();
    //private readonly string serverAddress;
    private readonly AuthModel authModel = new();

    public LoginViewModel()
    {
        SendAuthCommand = ReactiveCommand.Create(SendAuthAction);
        ServerPingCommand = ReactiveCommand.Create(ServerPingAction);
    }

    private async void ServerPingAction()
    {
        IsAuthInProcess = true;
        authModel.ChangeUrlParameters(serverField);
        try
        {
            await authModel.PingServer();
        }
        catch (Exception ex)
        {
            LogsOutputField = ex.ToString();
            IsAuthInProcess = true;
            return;
        }
        IsAuthInProcess = false;
    }

    private async void SendAuthAction()
    {
        IsAuthInProcess = true;
        user = new();
        try
        {
            if (IsAuth)
            {
                user = await authModel.GetUserInfoByAuth(new(loginField, passwordField));
            
            }
            else
            {
                user = await authModel.GetUserInfoByRegistration(new(loginField, nameField, passwordField));
            }
        }
        catch (Exception ex)
        {
            LogsOutputField = ex.ToString(); 
        }
        finally
        {

            LogsOutputField = JsonSerializer.Serialize(user);
            IsAuthInProcess = false;
        }

    }

    public async Task<UserInfo> GetUserInfoAsync(CancellationToken ct = default, int msDelay = 50)
    {
        while (!ct.IsCancellationRequested && user.auth_token == string.Empty)
        {
            await Task.Delay(msDelay);
        }
        return user;
    }
}
