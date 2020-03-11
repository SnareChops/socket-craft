export abstract class Droppable extends HTMLElement {
  protected init() {
    this.addEventListener('dragenter', (event: DragEvent) => this.onDragEnter(event));
    this.addEventListener('dragover', (event: DragEvent) => this.onDragOver(event));
    this.addEventListener('dragleave', (event: DragEvent) => this.onDragLeave(event));
  }

  protected onDragEnter(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
    this.classList.add('drop-highlight');
  }

  protected onDragOver(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
  }

  protected onDragLeave(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
    this.classList.remove('drop-highlight');
  }

  protected onDrop(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
    this.classList.remove('drop-highlight');
  }
}
