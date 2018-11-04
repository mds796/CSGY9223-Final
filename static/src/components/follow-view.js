import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class FollowView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <h2>Hello ${this.prop1}!</h2>
    `;
  }
  static get properties() {
    return {
      prop1: {
        type: String,
        value: 'follow-view'
      }
    };
  }
  constructor() {
    super();
    this.prop1 = 'Follow View';	  
  }
}

window.customElements.define('follow-view', FollowView);
