import { render, html, TemplateResult } from 'lit-html';
import { AuthService } from './services/AuthService';
import { getQueryParams } from './helpers';

export class SocketCraft extends HTMLElement {
  private get template(): TemplateResult {
    return html`
      <crafting-menu></crafting-menu>
      <div style="flex:1;"></div>
      <inventory-store></inventory-store>
    `;
  }

  private auth = new AuthService();

  public async connectedCallback() {
    const params = getQueryParams();
    if (!!params['token']) {
      // await this.auth.authenticate(params['token']);
      this.draw();
    }
    this.auth.login();
  }

  private draw() {
    render(this.template, this);
  }
}
customElements.define('socket-craft', SocketCraft);
