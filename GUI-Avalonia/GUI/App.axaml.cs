using Avalonia;
using Avalonia.Controls.ApplicationLifetimes;
using Avalonia.Markup.Xaml;

using GUI.ViewModels;
using GUI.Views;

namespace GUI;

public partial class App : Application
{
    public class Config
    {
        public string Server { get; set; }
    }

    public override void Initialize()
    {
        AvaloniaXamlLoader.Load(this);
    }

    public override async void OnFrameworkInitializationCompleted()
    {
/*        var configText = File.ReadAllText("config.json");
        Config config = JsonSerializer.Deserialize<Config>(configText);
*/
        if (ApplicationLifetime is IClassicDesktopStyleApplicationLifetime desktop)
        {
            var loginWindow = new LoginWindow()
            {
                DataContext = new LoginViewModel()
            };


            desktop.MainWindow = loginWindow;

            if (desktop.MainWindow.DataContext is LoginViewModel loginViewModel)
            {
                var userInfo = await loginViewModel.GetUserInfoAsync();
                //desktop.MainWindow.Close();
                var mainWindow = new MainWindow()
                {
                    DataContext = new MainViewModel(userInfo)
                };
                desktop.MainWindow = mainWindow;

                mainWindow.Show();
                loginWindow.Close();
            }
        }
        else if (ApplicationLifetime is ISingleViewApplicationLifetime singleViewPlatform)
        {
            var loginView = new LoginView()
            {
                DataContext = new LoginViewModel()
            };
            singleViewPlatform.MainView = loginView;

            if (singleViewPlatform.MainView.DataContext is LoginViewModel loginViewModel)
            {
                var userInfo = await loginViewModel.GetUserInfoAsync();
                
                var mainView = new MainView()
                {
                    DataContext = new MainViewModel(userInfo)
                };

                singleViewPlatform.MainView = mainView;

            }
        }

        base.OnFrameworkInitializationCompleted();

        /*        if (ApplicationLifetime is IClassicDesktopStyleApplicationLifetime desktop)
                {
                    desktop.MainWindow = new MainWindow
                    {
                        DataContext = new MainViewModel()
                    };
                }
                else if (ApplicationLifetime is ISingleViewApplicationLifetime singleViewPlatform)
                {
                    singleViewPlatform.MainView = new MainView
                    {
                        DataContext = new MainViewModel()
                    };
                }

                base.OnFrameworkInitializationCompleted();*/
    }
}
