import { html, render, TemplateResult } from 'lit-html';

export class InventoryStore extends HTMLElement {
  public get template(): TemplateResult {
    return html`
      <h2>Inventory</h2>
      <hr />
      <h3>Query Blocks</h3>
      <inventory-slots>
        <inventory-slot>
          <query-block></query-block>
        </inventory-slot>
      </inventory-slots>
    `;
  }

  public connectedCallback() {
    this.draw();
  }

  public draw() {
    render(this.template, this);
  }
}
customElements.define('inventory-store', InventoryStore);
