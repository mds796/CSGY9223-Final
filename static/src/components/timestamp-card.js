import {html, PolymerElement} from '@polymer/polymer/polymer-element.js';

/**
 * @customElement
 * @polymer
 */
class TimestampCard extends PolymerElement {
  static get template() {
    return html`
      <style>
        :host {
          display: block;
          border-radius: 25px;
          padding: 5px;
          margin-top: 5px;
          font-size: 0.7em;
          color: #777
        }
      </style>

      <span>[[timestamp]]@</span>
      <slot></slot>
    `;
  }

  static get properties() {
    return {
      timestamp: {
        type: String,
        value: ""
      }
    };
  }
}

window.customElements.define('timestamp-card', TimestampCard);
