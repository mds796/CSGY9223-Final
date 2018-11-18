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

      <span>Posted [[date]]</span>
      <slot></slot>
    `;
  }

  static get properties() {
    return {
      timestamp: {
        type: Number,
        value: 0
      },
      date: {
        type: String,
        computed: "dateString(timestamp)"
      }
    };
  }

  dateString(timestamp) {
    const date = new Date(0);
    date.setUTCSeconds(timestamp);
    return date
  }
}

window.customElements.define('timestamp-card', TimestampCard);
