import { html, render, TemplateResult } from 'lit-html';

export class CraftingMenu extends HTMLElement {
  public get template(): TemplateResult {
    return html`
      <h2>Crafting Menu</h2>
      <hr />
      <h3>Method Gems</h3>
      <crafting-drawer>
        <method-gem method="call">C</method-gem>
        <method-gem method="publish">P</method-gem>
        <method-gem method="subscribe">S</method-gem>
        <method-gem method="register">R</method-gem>
      </crafting-drawer>
    `;
  }

  public connectedCallback() {
    this.draw();
  }

  public draw() {
    render(this.template, this);
  }
}
customElements.define('crafting-menu', CraftingMenu);
