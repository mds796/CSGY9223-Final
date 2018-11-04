import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class AboutView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <h2>About view!</h2>
    `;
  }
}

window.customElements.define('about-view', AboutView);
