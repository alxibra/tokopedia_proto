# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: access_token.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("access_token.proto", :syntax => :proto3) do
    add_message "tokopedia.AccessToken" do
      optional :access_token, :string, 1
      optional :expires_in, :uint32, 2
    end
  end
end

module Tokopedia
  AccessToken = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("tokopedia.AccessToken").msgclass
end
