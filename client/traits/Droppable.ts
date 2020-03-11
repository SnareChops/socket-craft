import { prependConnected } from './helpers';

export function Droppable(): ClassDecorator {
  return prependConnected(function(this: HTMLElement) {
    this.addEventListener('dragenter', (event: DragEvent) => {
      event.preventDefault();
      event.stopPropagation();
      this.classList.add('drop-highlight');
    });

    this.addEventListener('dragover', (event: DragEvent) => {
      event.preventDefault();
      event.stopPropagation();
    });

    this.addEventListener('dragleave', (event: DragEvent) => {
      event.preventDefault();
      event.stopPropagation();
      this.classList.remove('drop-highlight');
    });

    this.addEventListener('drop', (event: DragEvent) => {
      event.stopPropagation();
      this.classList.remove('drop-highlight');
      if ((this as any).onDrop && (this as any).onDrop(event)) {
        event.preventDefault();
      }
    });
  });
}
