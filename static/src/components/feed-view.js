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

            textarea {
              padding: 10px;
              resize: none;
            }
          </style>

          <section class="post-input">
            <textarea placeholder="What's happening?" value="{{post::change}}" rows="7" cols="80">
            </textarea>
            <input id="post-button" type="submit" value="Post" on-click="submitPost"/>
          </section>

          <dom-repeat items="[[feed]]">
            <template>
              <section>
                <post-card user="[[item.User.Name]]" timestamp="[[item.Timestamp.EpochNanoseconds]]" text="[[item.Text]]"></post-card>
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

    submitPost() {
        const provider = this;
        const submitButton = this.shadowRoot.getElementById('post-button');

        submitButton.disabled = true;

        fetch('/post', {method: 'POST', body: provider.post, headers: {'Content-Type': 'text/plain'}})
            .then(response => {
                if (response.ok) {
                    provider.post = "";
                }

                submitButton.disabled = false;
                provider.fetchFeed();
            })
            .catch(err => {
                this.shadowRoot.getElementById('post-button').disabled = false;
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
            console.log("Unable to fetch feed: ", err);
        });

        setTimeout(() => this.fetchFeed(), 1000);
    }
}

window.customElements.define('feed-view', FeedView);
