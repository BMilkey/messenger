

using System.Globalization;
using System.Reflection;



namespace TestBox;

public static class Program
{
    public static void Main(string[] args)
    {
        DateTime dateTime = DateTime.Now;
        Console.WriteLine(dateTime);
        Console.WriteLine(dateTime.ToString());
        Console.WriteLine(dateTime.ToShortDateString());
        Console.WriteLine(dateTime.ToShortTimeString());
        Console.WriteLine(dateTime.ToLongDateString());
        Console.WriteLine(dateTime.ToLongTimeString());
        Console.WriteLine();
        Timestamp.TimestampAttribute timestamp = new(dateTime.ToString());
        Console.WriteLine(timestamp.Timestamp);

        Console.WriteLine(RetrieveTimestamp());
        Console.WriteLine(RetrieveTimestampAsDateTime());

    }

    public static string RetrieveTimestamp()
    {
        var attribute = Assembly.GetExecutingAssembly()
            .GetCustomAttributesData()
            .First(x => x.AttributeType.Name == "TimestampAttribute");

        return (string)attribute.ConstructorArguments.First().Value;
    }

    public static DateTime RetrieveTimestampAsDateTime()
    {
        var timestamp = RetrieveTimestamp();
        return DateTime.ParseExact(timestamp, "yyyy-MM-ddTHH:mm:ss.fffZ", null, DateTimeStyles.AssumeUniversal)
            .ToUniversalTime();
    }
}
