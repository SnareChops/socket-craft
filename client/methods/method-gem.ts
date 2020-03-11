import { render, html } from 'lit-html';
import { Draggable } from '../traits/Draggable';

@Draggable()
export class MethodGem extends HTMLElement {
  public get method(): string {
    return this.getAttribute('method');
  }

  public get dataTransfer(): { [key: string]: string } {
    return {
      type: 'method-gem',
      id: this.method
    };
  }

  public connectedCallback() {
    // prettier-ignore
    render(html`${this.method[0].toUpperCase()}`, this);
  }
}
customElements.define('method-gem', MethodGem);
