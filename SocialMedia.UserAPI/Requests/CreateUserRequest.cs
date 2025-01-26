namespace SocialMedia.UserAPI.Requests;

public class CreateUserRequest
{
    public string Username { get; set; }
    public string DisplayName { get; set; }
    public string? Biography { get; set; }
    public string? Location { get; set; }
    public string? Website { get; set; }
}