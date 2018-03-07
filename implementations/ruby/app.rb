require 'dotenv/load'
require 'faraday'
require 'json'
require 'sinatra'

require_relative 'messaging_handler'

configure do
  token = ENV['ACCESS_TOKEN']
  fb_url = "https://graph.facebook.com/v2.6/me/messages?access_token=#{token}"
  fb_client = Faraday.new(url: fb_url) do |conn|
    conn.request :url_encoded
    conn.adapter Faraday.default_adapter
    conn.response :logger
    conn.headers['Content-Type'] = 'application/json'
  end

  set :fb_client, fb_client
  set :verification_token, ENV['VERIFICATION_TOKEN']
end

get '/webhook' do
  mode = params['hub.mode']
  token = params['hub.verify_token']
  challenge = params['hub.challenge']

  return challenge if mode == 'subscribe' && token == settings.verification_token

  halt 403
end

post '/webhook' do
  request.body.rewind
  data = JSON.parse(request.body.read)

  data['entry'].each do |entry|
    entry['messaging'].each do |messaging|
      process_messaging(messaging)
    end
  end

  200
end

private

def process_messaging(messaging)
  sender_id = messaging.dig('sender', 'id')

  if messaging.key? 'message'
    handler = TextMessageHandler.new(
      sender_id: sender_id,
      client: settings.fb_client
    )
    handler.handle(messaging)
  elsif messaging.key? 'postback'
    handler = PostBackHandler.new(
      sender_id: sender_id,
      client: settings.fb_client
    )
    handler.handle(messaging['postback'])
  end
end
