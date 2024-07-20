using GUI.Models.Main;
using ReactiveUI;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reactive;
using System.Text;
using System.Threading.Tasks;

namespace GUI.ViewModels.Main;

public class CurrentChatViewModel : ViewModelBase
{
    #region xaml fields

    public ReactiveCommand<string, Unit> CreateMessageCommand { get; set; }
    #endregion

    private readonly MainModel mainModel;
    public MainModel MainModel => mainModel;

    public CurrentChatViewModel(MainModel mainModel)
    {
        this.mainModel = mainModel;
        CreateMessageCommand = ReactiveCommand.Create<string>(CreateMessageAction);
    }

    public CurrentChatViewModel()
    {

    }

    public void CreateMessageAction(string message)
    {
        _ = MainModel.CreateMessage(message);
        
    }
}
