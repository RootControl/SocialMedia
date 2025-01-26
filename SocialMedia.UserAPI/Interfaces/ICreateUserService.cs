using SocialMedia.UserAPI.Requests;
using SocialMedia.UserAPI.Responses;

namespace SocialMedia.UserAPI.Interfaces;

public interface ICreateUserService
{
    public Task<CreateUserResponse> CreateUserAsync(CreateUserRequest request);
}