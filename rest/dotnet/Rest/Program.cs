using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using System;
using System.Collections.Generic;
using System.Linq;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Http.Json;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.WebHost.ConfigureKestrel(options =>
{
    options.ListenAnyIP(8080); // Set your desired port here
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.MapPost("/compute", async (HttpContext context) =>
{
    try
    {
        var numbers = await context.Request.ReadFromJsonAsync<Numbers>();
        
        if (numbers == null || !numbers.Values.Any())
            return Results.BadRequest("Invalid input.");

        var sum = numbers.Values.Sum();
        var product = numbers.Values.Aggregate(1.0, (acc, val) => acc * val);
        var average = sum / numbers.Values.Count;

        return Results.Ok(new Result
        {
            Sum = sum,
            Average = average,
            Product = product
        });
    }
    catch (Exception ex)
    {
        return Results.BadRequest(ex.Message);
    }
});


app.Run();

record Numbers(List<double> Values);

record Result
{
    public double Sum { get; init; }
    public double Average { get; init; }
    public double Product { get; init; }
};
