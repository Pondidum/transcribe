12 Factor Apps send all their logs as structured lines (usually in json) to stdout for another process to deal with.  This is great for keeping the app simple, but it makes the logs hard to read locally.

Enter `transcribe.exe`, which parses Serilog's output and renders it to the console for you.

# Example Output

Just running the `LogLines.dll` shows this serilog output:

```bash
$ dotnet LogLines.dll
{"Timestamp":"2018-08-20T17:47:35.5366638+03:00","Level":"Information","MessageTemplate":"This is the {count} message","Properties":{"count":1}}
{"Timestamp":"2018-08-20T17:47:36.5481673+03:00","Level":"Information","MessageTemplate":"This is the {count} message","Properties":{"count":2}}
{"Timestamp":"2018-08-20T17:47:37.5495459+03:00","Level":"Information","MessageTemplate":"This is the {count} message","Properties":{"count":3}}
```

Just append `| transcribe` and we get this output instead:

```bash
$ dotnet LogLines.dll | transcribe
This is the 1 message
This is the 2 message
This is the 3 message
```

# C# Setup

You need the following NuGet packages:

```bash
dotnet add package Serilog
dotnet add package Serilog.Sinks.Console
```

static void Main(string[] args)
{
	Log.Logger = new LoggerConfiguration()
		.WriteTo.Console(new JsonFormatter())
		.CreateLogger();

	var i = 0;
	while (true)
	{
		i++;
		Log.Information("This is the {count} message", i);
		Thread.Sleep(1000);
	}
}

# Contributing

I have no idea what I am doing with Go really.  Any help appreciated.

# Todo List

* Coloured output?
* Handling complex objects
* handle Serilog Json compressed format?
* investigate why gitbash and vscode consoles can't pipe dotnet.exe output (to anything)
