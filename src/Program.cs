using Newtonsoft.Json;


var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment()) {
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.MapGet("/words", async (HttpRequest request, HttpResponse response) => {
    var now      = DateTime.Now;
    var tomorrow = now.AddDays(1).Date;
    string lang  = request.Query["lang"];

    IWordleRepository? repository = lang switch {
        "ru" => WordleRepository<RussianWord>.GetRussian(),
        "en" => WordleRepository<EnglishWord>.GetEnglish(),
        _    => null
    };

    if (repository == null) {
        response.StatusCode = StatusCodes.Status400BadRequest;
        await response.WriteAsync("{\"error\": \"unknown language\"}");
    } else {
        await response.WriteAsync(JsonConvert.SerializeObject(new Payload {
            lang       = lang,
            word       = repository.Sync().GetWordOfDay()?.Word ?? "",
            timeToNext = (UInt64)(tomorrow - now).TotalSeconds
        }));
    }
});

app.Run();

