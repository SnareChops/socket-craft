import { prependConnected } from './helpers';

export function Draggable(): ClassDecorator {
  return prependConnected(function(this: HTMLElement) {
    this.setAttribute('draggable', 'true');
    this.addEventListener('dragstart', (event: DragEvent) => {
      if (!!(this as any).dataTransfer) {
        for (const key in (this as any).dataTransfer) {
          console.log('setting dataTransfer', key, (this as any).dataTransfer[key]);
          event.dataTransfer.setData(key, (this as any).dataTransfer[key]);
        }
      }
    });
  });
}
