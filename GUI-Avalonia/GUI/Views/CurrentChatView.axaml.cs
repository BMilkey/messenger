using Avalonia.Controls;

namespace GUI.Views
{
    public partial class CurrentChatView : UserControl
    {
        public CurrentChatView()
        {
            InitializeComponent();
        }

        private void TextBox_KeyDown(object? sender, Avalonia.Input.KeyEventArgs e)
        {
            if (e.Key == Avalonia.Input.Key.Enter)
            {
                var tb = sender as TextBox;
                this.SendMessageButton.Command.Execute(tb.Text);
            }
        }
    }
}
