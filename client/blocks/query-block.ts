import { render, html, TemplateResult } from 'lit-html';

export class QueryBlock extends HTMLElement {
  public get template(): TemplateResult {
    return html`
      <method-socket></method-socket>
    `;
  }

  public connectedCallback() {
    this.draw();
  }

  public draw() {
    render(this.template, this);
  }
}
customElements.define('query-block', QueryBlock);
