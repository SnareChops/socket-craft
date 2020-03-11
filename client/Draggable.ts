export abstract class Draggable extends HTMLElement {
  protected init() {
    this.setAttribute('draggable', 'true');
    this.addEventListener('dragstart', (event: DragEvent) => this.dragstart(event));
  }

  protected dragstart(event: DragEvent) {
    console.log('dragstart', this);
    event.stopPropagation();
    event.dataTransfer.setData('text/', '');
    event.dataTransfer.effectAllowed = 'copy';
  }
}
