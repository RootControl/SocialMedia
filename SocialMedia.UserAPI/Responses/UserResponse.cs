namespace SocialMedia.UserAPI.Responses;

public abstract class UserResponse
{
    public bool Success { get; protected set; }
    public string? ErrorMessage { get; protected set; }
}