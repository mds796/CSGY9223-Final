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
      user: { type: Object },
      feed: { type: Array },
      influencers: { type: Array }
    };
  }
}

window.customElements.define('data-provider', DataProvider);
