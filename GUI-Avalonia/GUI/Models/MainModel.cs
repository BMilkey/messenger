using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace GUI.Models;

public class MainModel
{
    private HttpClient client;
    private string urlParameters;

    public MainModel(string URL, string urlParameters = "")
    {
        this.urlParameters = urlParameters;
        client = new HttpClient();
    }

    
}
