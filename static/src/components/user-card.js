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
          background: #cccccc;
          border-radius: 25px;
          padding: 5px;
        }
      </style>

      <span>[[user]]@</span>
      <slot></slot>
    `;
  }

  static get properties() {
    return {
      user: {
        type: String,
        value: ""
      }
    };
  }
}

window.customElements.define('user-card', UserCard);
