import { render, html, TemplateResult } from 'lit-html';
import { Droppable } from '../traits/Droppable';

@Droppable()
export class MethodSocket extends HTMLElement {
  public method: 'call' | 'publish' | 'subscribe' | 'register';

  public get template(): TemplateResult {
    return html`
      ${!!this.method
        ? html`
            <method-gem method="${this.method}"></method-gem>
          `
        : ''}
    `;
  }

  public connectedCallback() {
    this.draw();
  }

  public draw() {
    render(this.template, this);
  }

  public onDrop(event: DragEvent): boolean {
    if (event.dataTransfer.getData('type') !== 'method-gem') {
      return false;
    }
    this.method = void 0;
    this.draw();
    const id = event.dataTransfer.getData('id');
    this.method = id as any;
    this.draw();
    return true;
  }
}
customElements.define('method-socket', MethodSocket);
