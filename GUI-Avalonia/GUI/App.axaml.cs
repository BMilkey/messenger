using Avalonia;
using Avalonia.Controls.ApplicationLifetimes;
using Avalonia.Markup.Xaml;

using GUI.ViewModels;
using GUI.Views;
using System;
using System.Diagnostics;

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
                var (userInfo, URL) = await loginViewModel.GetUserInfoAndURLAsync();
                //desktop.MainWindow.Close();
                var mainWindow = new MainWindow()
                {
                    DataContext = new MainViewModel(userInfo, URL)
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
                var (userInfo, URL) = await loginViewModel.GetUserInfoAndURLAsync();
                
                var mainView = new MainView()
                {
                    DataContext = new MainViewModel(userInfo, URL)
                };

                singleViewPlatform.MainView = mainView;
                Console.WriteLine(singleViewPlatform.MainView);
                Debug.WriteLine(singleViewPlatform.MainView);
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
