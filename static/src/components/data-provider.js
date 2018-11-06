import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';

/**
 * @customElement
 * @polymer
 */
class DataProvider extends PolymerElement {
  static get template() {
    return html`
      <style>
        :host {
          display: hidden;
        }
      </style>
    `;
  }
  
  static get properties() {
    return {
      user: { type: Object, value: {} },
      posts: { type: Array, value: [] },
      influencers: { type: Array, value: [] }
    };
  }
}

window.customElements.define('data-provider', DataProvider);
