using SocialMedia.UserAPI.Interfaces;
using SocialMedia.UserAPI.Requests;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

app.UseHttpsRedirection();

app.MapPost("/user", async (ICreateUserService createUserService, CreateUserRequest request) =>
{
    var response = await createUserService.CreateUserAsync(request);
    
    if (!response.Success)
        return Results.BadRequest(response);
    
    return Results.Created($"/user/{response.User?.Username}", response);
});

app.Run();