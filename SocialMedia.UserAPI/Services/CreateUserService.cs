using SocialMedia.UserAPI.Entities;
using SocialMedia.UserAPI.Interfaces;
using SocialMedia.UserAPI.Requests;
using SocialMedia.UserAPI.Responses;

namespace SocialMedia.UserAPI.Services;

public class CreateUserService : ICreateUserService
{
    public async Task<CreateUserResponse> CreateUserAsync(CreateUserRequest request)
    {
        if (!RequestIsValid(request))
            return new CreateUserResponse();
        
        var user = new User(request);
        
        // TODO: Save user to database

        return new CreateUserResponse(user);
    }
    
    private bool RequestIsValid(CreateUserRequest request)
    {
        return !string.IsNullOrWhiteSpace(request.Username) && 
               !string.IsNullOrWhiteSpace(request.DisplayName);
    }
}