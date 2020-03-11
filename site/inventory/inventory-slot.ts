import { Droppable } from '../traits/Droppable';

@Droppable()
export class InventorySlot extends HTMLElement {}
customElements.define('inventory-slot', InventorySlot);
