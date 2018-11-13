import {html, PolymerElement} from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

import './post-card.js';

/**
 * @customElement
 * @polymer
 */
class FeedView extends PolymerElement {
    static get template() {
        return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <style>
            :host .post-input {
              text-align: center;
            }
          </style>

          <section class="post-input">
            <textarea value="{{post::change}}" rows="7" cols="80">
            </textarea>
            <input type="submit" value="Post" on-click="submitPost"/>
          </section>

          <dom-repeat items="[[feed]]">
            <template>
              <section>
                <post-card user="[[item.From]]" timestamp="[[item.Timestamp]]" text="[[item.Text]]"></post-card>
              </section>
            </template>
          </dom-repeat>
        </template>
      </dom-if>
    `;
    }

    static get properties() {
        return {
            active: {type: Boolean, value: false, observer: 'fetchFeed'},
            post: {type: String, value: ""},
            feed: {type: Array, value: []}
        };
    }

    ready() {
        super.ready();
    }

    submitPost(e) {
        const provider = this;

        e.target.disabled = true;

        fetch('/post', {method: 'POST', body: provider.post, headers: {'Content-Type': 'text/plain'}})
            .then(_ => {
                provider.post = "";
                e.target.disabled = false;
                provider.fetchFeed();
            })
            .catch(err => {
                e.target.disabled = false;
                provider.fetchFeed();
                console.log("Unable to post new message: ", err);
            });
    }

    fetchFeed() {
        const provider = this;

        if (!provider.active) {
            return;
        }

        fetch('/feed').then(response => {
            if (response.ok) {
                return response.json();
            } else {
                return {};
            }
        }).then(data => {
            provider.feed = data.Posts;
        }).catch(err => {
            console.log("Unable to fetch follows: ", err);
        });
    }
}

window.customElements.define('feed-view', FeedView);
