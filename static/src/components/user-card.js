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
      
      <span>[[name]]@</span>
      <slot></slot>
    `;
  }

  static get properties() {
    return {
      name: {
        type: String,
        value: ""
      }
    };
  }
}

window.customElements.define('user-card', UserCard);
