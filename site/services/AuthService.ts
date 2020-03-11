export class AuthService {
  private get clientId() {
    return 'e2ba2e865e1196d2e044';
  }

  private get redirectUri() {
    return 'http://localhost:8080';
  }

  private get state() {
    return 'fsalkhqefgasihdf';
  }

  private get oauthUrl(): string {
    return [
      `https://github.com/login/oauth/authorize`,
      `?client_id=${this.clientId}`,
      `&redirect_uri=${this.redirectUri}`,
      `&state=${this.state}`
    ].join('');
  }

  public login() {
    location.href = this.oauthUrl;
  }

  public authorize(token: string, state: string) {
    if (state !== this.state) {
      throw new Error('State Mismatch');
    }
  }
}
