using SocialMedia.UserAPI.Requests;

namespace SocialMedia.UserAPI.Entities;

public class User
{
    public string Username { get; private set; }
    public string DisplayName { get; private set; }
    public string? Biography { get; private set; }
    public string? Location { get; private set; }
    public string? Website { get; private set; }
    public DateTime JoinedDate { get; private set; }

    public User(CreateUserRequest request)
    {
        Username = request.Username;
        DisplayName = request.DisplayName;
        Biography = request.Biography;
        Location = request.Location;
        Website = request.Website;
        JoinedDate = DateTime.Now;
    }
}