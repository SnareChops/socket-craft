export function prependConnected(prepended: (this: HTMLElement) => void): (target: any) => void {
  return target => {
    const original = target.prototype.connectedCallback;
    target.prototype.connectedCallback = function(this: HTMLElement) {
      prepended && prepended.bind(this)();
      original && original.bind(this)();
    };
    return target;
  };
}
