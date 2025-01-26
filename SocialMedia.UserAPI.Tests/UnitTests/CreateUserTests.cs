using SocialMedia.UserAPI.Requests;
using SocialMedia.UserAPI.Services;

namespace SocialMedia.UserAPI.Tests.UnitTests;

public class CreateUserTests
{
    private readonly VerifySettings _verifySettings;

    public CreateUserTests()
    {
        _verifySettings = new VerifySettings();
        _verifySettings.UseDirectory("snapshots");
    }
        
    [Fact]
    public Task CreateUserAsync_ShouldReturnError_WhenValidatingRequest()
    {
        var request = new CreateUserRequest();
        var service = new CreateUserService();
        
        var response = service.CreateUserAsync(request);
        
        return Verify(response, _verifySettings);
    }
    
    [Fact]
    public Task CreateUserAsync_ShouldReturnError_WhenRequestIsInvalid()
    {
        var request = new CreateUserRequest
        {
            Username = "johndoe"
        };
        var service = new CreateUserService();
        
        var response = service.CreateUserAsync(request);
        
        return Verify(response, _verifySettings);
    }
    
    [Fact]
    public Task CreateUserAsync_ShouldReturnSuccess_WhenRequestIsValid()
    {
        var request = new CreateUserRequest
        {
            Username = "johndoe",
            DisplayName = "John Doe"
        };
        var service = new CreateUserService();
        
        var response = service.CreateUserAsync(request);
        
        return Verify(response, _verifySettings);
    }
}