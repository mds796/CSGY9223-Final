import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class NotFoundView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <section>
        <h2>Oops! You hit a 404</h2>
        <p>The page you're looking for doesn't seem to exist. Head back
           <a href="/">home</a> and try again?
        </p>
      </section>
    `;
  }
}

window.customElements.define('not-found-view', NotFoundView);
