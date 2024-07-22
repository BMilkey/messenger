using AsyncImageLoader.Loaders;
using AsyncImageLoader;
using GUI.Models.Main;
using ReactiveUI;
using System.Windows.Input;
using GUI.Models;
using System.Collections.ObjectModel;
using System.Reactive;
using System.Diagnostics;

namespace GUI.ViewModels.Main;

public class ChatsViewModel : ViewModelBase
{
    #region xaml fields
    private IAsyncImageLoader imageLoader = new RamCachedWebImageLoader();
    public IAsyncImageLoader ImageLoader => imageLoader;

    public int MainListWidth => 400 + (isCreateChatToggled ? 200 : 0);

    private string findUserOrChatField = string.Empty;
    public string FindUserOrChatField
    {
        get => findUserOrChatField;
        set
        {
            this.RaiseAndSetIfChanged(ref findUserOrChatField, value);
            _ = mainModel.FindUsersByName(value);
        }
    }

    public string ToggleCreateChatButtonContent
    {
        get => isCreateChatToggled ? ">" : "<";
    }

    private bool isCreateChatToggled = true;
    public bool IsCreateChatToggled
    {
        get => isCreateChatToggled;
        set => this.RaiseAndSetIfChanged(ref isCreateChatToggled, value);
    }


    public ReactiveCommand<UserInfo, Unit> RemoveUserFromSelectedCommand { get; set; }
    public ReactiveCommand<UserInfo, Unit> AddUserToCreateChatListCommand { get; set; }
    public ReactiveCommand<string, Unit> CreateChatCommand { get; set; }

    private ObservableCollection<UserInfo> newChatUsers = new();
    public ObservableCollection<UserInfo> NewChatUsers
    {
        get => newChatUsers;
        set => this.RaiseAndSetIfChanged(ref newChatUsers, value);
    }
    #endregion

    private readonly MainModel mainModel;
    public MainModel MainModel => mainModel;


    public ChatsViewModel(MainModel mainModel)
    {
        this.mainModel = mainModel;
        AddUserToCreateChatListCommand = ReactiveCommand.Create<UserInfo>(AddUserToCreateChatListAction);
        CreateChatCommand = ReactiveCommand.Create<string>(CreateChatAction);
        RemoveUserFromSelectedCommand = ReactiveCommand.Create<UserInfo>(RemoveUserFromSelectedAction);
    }

    public ChatsViewModel()
    {

    }

    public void RemoveUserFromSelectedAction(UserInfo userInfo)
    {
        newChatUsers.Remove(userInfo);
    }

    public void AddUserToCreateChatListAction(UserInfo userInfo)
    {
        if (newChatUsers.Contains(userInfo))
            return;

        newChatUsers.Add(userInfo);
    }

    public void CreateChatAction(string title)
    {
        // TODO
        // clear on success or smth else
        _ = mainModel.CreateChat(title, newChatUsers);
    }


}
