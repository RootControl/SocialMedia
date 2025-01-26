using SocialMedia.UserAPI.Entities;

namespace SocialMedia.UserAPI.Responses;

public class CreateUserResponse : UserResponse
{
    public User? User { get; private set; }
    
    public CreateUserResponse(User? user = null)
    {
        User = user;
        Success = user != null;
        ErrorMessage = user == null ? "Failed to create user" : null;
    }
}