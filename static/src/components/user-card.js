import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';

/**
 * @customElement
 * @polymer
 */
class UserCard extends PolymerElement {
  static get template() {
    return html`
      <style>
        :host {
          display: block;
        }
      </style>
      
      <span>[[user]]@</span>
      <slot></slot>
    `;
  }

  static get properties() {
    return {
      user: {
        type: String
      }
    };
  }
}

window.customElements.define('user-card', UserCard);
