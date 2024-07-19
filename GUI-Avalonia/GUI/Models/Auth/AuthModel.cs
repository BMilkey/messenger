using System;
using System.Diagnostics;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Security.Cryptography.X509Certificates;
using System.Threading.Tasks;

namespace GUI.Models.Auth;

public class AuthModel
{

    public class AuthRequest
    {
        public string Login { get; set; } = string.Empty;
        public string Password { get; set; } = string.Empty;
        public AuthRequest(string login, string password)
        {
            Login = login;
            Password = password;
        }
    }

    public class RegistrationRequest
    {
        public string Login { get; set; } = string.Empty;
        public string Name { get; set; } = string.Empty;
        public string Password { get; set; } = string.Empty;
        public RegistrationRequest(string login, string name, string password)
        {
            Login = login;
            Name = name;
            Password = password;
        }
    }

    private HttpClient client;
    private string urlParameters;
    public string URL { get; private set; }

    public AuthModel(string URL, string urlParameters = "")
    {
        this.URL = URL;
        this.urlParameters = urlParameters;
        client = new HttpClient();
        client.BaseAddress = new Uri(URL);
    }

    public AuthModel()
    {
        urlParameters = "";
        client = new HttpClient();
    }

    public void ChangeUrlParameters(string URL, string urlParameters = "")
    {
        this.URL = URL;
        client = new HttpClient();
        client.BaseAddress = new Uri($"http://{URL}/");
        this.urlParameters = urlParameters;

    }

    public async Task PingServer()
    {
        var response = await client.PostAsync(
                                "/post/test/ping", null);
        response.EnsureSuccessStatusCode();
    }

    public async Task<UserInfo> GetUserInfoByAuth(AuthRequest request)
    {

        client.DefaultRequestHeaders.Accept.Add(
                new MediaTypeWithQualityHeaderValue("application/json"));

        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/auth/user_by_auth", request);
        response.EnsureSuccessStatusCode();

        var registerAnswer = new UserInfo();

        /*        var stringAnswer = await response.Content.ReadFromJsonAsync<string>();
                Debug.WriteLine(stringAnswer);
                var streamAnswer = await response.Content.ReadAsStreamAsync();
                Debug.WriteLine(streamAnswer);*/

        registerAnswer = await response.Content.ReadFromJsonAsync<UserInfo>();

        return registerAnswer;
    }
    public async Task<UserInfo> GetUserInfoByRegistration(RegistrationRequest request)
    {

        // Add an Accept header for JSON format.
        client.DefaultRequestHeaders.Accept.Add(
                new MediaTypeWithQualityHeaderValue("application/json"));

        // List data response.
        HttpResponseMessage response = await client.PostAsJsonAsync(
                                            "/post/auth/register_user", request);
        response.EnsureSuccessStatusCode();

        var registerAnswer = new UserInfo();



        // Parse the response body.
        var stringAnswer = await response.Content.ReadAsStreamAsync();
        Debug.WriteLine(stringAnswer);

        registerAnswer = await response.Content.ReadFromJsonAsync<UserInfo>();  //Make sure to add a reference to System.Net.Http.Formatting.dll

        Debug.WriteLine($"parsed answer: {registerAnswer}");



        return registerAnswer;

    }

}
