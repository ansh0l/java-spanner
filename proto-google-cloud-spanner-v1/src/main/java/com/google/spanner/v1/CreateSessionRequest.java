/*
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: google/spanner/v1/spanner.proto

package com.google.spanner.v1;

/**
 *
 *
 * <pre>
 * The request for [CreateSession][google.spanner.v1.Spanner.CreateSession].
 * </pre>
 *
 * Protobuf type {@code google.spanner.v1.CreateSessionRequest}
 */
public final class CreateSessionRequest extends com.google.protobuf.GeneratedMessageV3
    implements
    // @@protoc_insertion_point(message_implements:google.spanner.v1.CreateSessionRequest)
    CreateSessionRequestOrBuilder {
  private static final long serialVersionUID = 0L;
  // Use CreateSessionRequest.newBuilder() to construct.
  private CreateSessionRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }

  private CreateSessionRequest() {
    database_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(UnusedPrivateParameter unused) {
    return new CreateSessionRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet getUnknownFields() {
    return this.unknownFields;
  }

  private CreateSessionRequest(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10:
            {
              java.lang.String s = input.readStringRequireUtf8();

              database_ = s;
              break;
            }
          case 18:
            {
              com.google.spanner.v1.Session.Builder subBuilder = null;
              if (session_ != null) {
                subBuilder = session_.toBuilder();
              }
              session_ =
                  input.readMessage(com.google.spanner.v1.Session.parser(), extensionRegistry);
              if (subBuilder != null) {
                subBuilder.mergeFrom(session_);
                session_ = subBuilder.buildPartial();
              }

              break;
            }
          default:
            {
              if (!parseUnknownField(input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }

  public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
    return com.google.spanner.v1.SpannerProto
        .internal_static_google_spanner_v1_CreateSessionRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.google.spanner.v1.SpannerProto
        .internal_static_google_spanner_v1_CreateSessionRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.google.spanner.v1.CreateSessionRequest.class,
            com.google.spanner.v1.CreateSessionRequest.Builder.class);
  }

  public static final int DATABASE_FIELD_NUMBER = 1;
  private volatile java.lang.Object database_;
  /**
   *
   *
   * <pre>
   * Required. The database in which the new session is created.
   * </pre>
   *
   * <code>
   * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
   * </code>
   *
   * @return The database.
   */
  @java.lang.Override
  public java.lang.String getDatabase() {
    java.lang.Object ref = database_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      database_ = s;
      return s;
    }
  }
  /**
   *
   *
   * <pre>
   * Required. The database in which the new session is created.
   * </pre>
   *
   * <code>
   * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
   * </code>
   *
   * @return The bytes for database.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getDatabaseBytes() {
    java.lang.Object ref = database_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b =
          com.google.protobuf.ByteString.copyFromUtf8((java.lang.String) ref);
      database_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int SESSION_FIELD_NUMBER = 2;
  private com.google.spanner.v1.Session session_;
  /**
   *
   *
   * <pre>
   * The session to create.
   * </pre>
   *
   * <code>.google.spanner.v1.Session session = 2;</code>
   *
   * @return Whether the session field is set.
   */
  @java.lang.Override
  public boolean hasSession() {
    return session_ != null;
  }
  /**
   *
   *
   * <pre>
   * The session to create.
   * </pre>
   *
   * <code>.google.spanner.v1.Session session = 2;</code>
   *
   * @return The session.
   */
  @java.lang.Override
  public com.google.spanner.v1.Session getSession() {
    return session_ == null ? com.google.spanner.v1.Session.getDefaultInstance() : session_;
  }
  /**
   *
   *
   * <pre>
   * The session to create.
   * </pre>
   *
   * <code>.google.spanner.v1.Session session = 2;</code>
   */
  @java.lang.Override
  public com.google.spanner.v1.SessionOrBuilder getSessionOrBuilder() {
    return getSession();
  }

  private byte memoizedIsInitialized = -1;

  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output) throws java.io.IOException {
    if (!getDatabaseBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, database_);
    }
    if (session_ != null) {
      output.writeMessage(2, getSession());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getDatabaseBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, database_);
    }
    if (session_ != null) {
      size += com.google.protobuf.CodedOutputStream.computeMessageSize(2, getSession());
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
      return true;
    }
    if (!(obj instanceof com.google.spanner.v1.CreateSessionRequest)) {
      return super.equals(obj);
    }
    com.google.spanner.v1.CreateSessionRequest other =
        (com.google.spanner.v1.CreateSessionRequest) obj;

    if (!getDatabase().equals(other.getDatabase())) return false;
    if (hasSession() != other.hasSession()) return false;
    if (hasSession()) {
      if (!getSession().equals(other.getSession())) return false;
    }
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + DATABASE_FIELD_NUMBER;
    hash = (53 * hash) + getDatabase().hashCode();
    if (hasSession()) {
      hash = (37 * hash) + SESSION_FIELD_NUMBER;
      hash = (53 * hash) + getSession().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      java.nio.ByteBuffer data, com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      byte[] data, com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseWithIOException(PARSER, input);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      java.io.InputStream input, com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseWithIOException(
        PARSER, input, extensionRegistry);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseDelimitedFrom(
      java.io.InputStream input) throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseDelimitedWithIOException(PARSER, input);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseDelimitedFrom(
      java.io.InputStream input, com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseDelimitedWithIOException(
        PARSER, input, extensionRegistry);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      com.google.protobuf.CodedInputStream input) throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseWithIOException(PARSER, input);
  }

  public static com.google.spanner.v1.CreateSessionRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3.parseWithIOException(
        PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() {
    return newBuilder();
  }

  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }

  public static Builder newBuilder(com.google.spanner.v1.CreateSessionRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }

  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   *
   *
   * <pre>
   * The request for [CreateSession][google.spanner.v1.Spanner.CreateSession].
   * </pre>
   *
   * Protobuf type {@code google.spanner.v1.CreateSessionRequest}
   */
  public static final class Builder extends com.google.protobuf.GeneratedMessageV3.Builder<Builder>
      implements
      // @@protoc_insertion_point(builder_implements:google.spanner.v1.CreateSessionRequest)
      com.google.spanner.v1.CreateSessionRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor getDescriptor() {
      return com.google.spanner.v1.SpannerProto
          .internal_static_google_spanner_v1_CreateSessionRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.google.spanner.v1.SpannerProto
          .internal_static_google_spanner_v1_CreateSessionRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.google.spanner.v1.CreateSessionRequest.class,
              com.google.spanner.v1.CreateSessionRequest.Builder.class);
    }

    // Construct using com.google.spanner.v1.CreateSessionRequest.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }

    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders) {}
    }

    @java.lang.Override
    public Builder clear() {
      super.clear();
      database_ = "";

      if (sessionBuilder_ == null) {
        session_ = null;
      } else {
        session_ = null;
        sessionBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor getDescriptorForType() {
      return com.google.spanner.v1.SpannerProto
          .internal_static_google_spanner_v1_CreateSessionRequest_descriptor;
    }

    @java.lang.Override
    public com.google.spanner.v1.CreateSessionRequest getDefaultInstanceForType() {
      return com.google.spanner.v1.CreateSessionRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.google.spanner.v1.CreateSessionRequest build() {
      com.google.spanner.v1.CreateSessionRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.google.spanner.v1.CreateSessionRequest buildPartial() {
      com.google.spanner.v1.CreateSessionRequest result =
          new com.google.spanner.v1.CreateSessionRequest(this);
      result.database_ = database_;
      if (sessionBuilder_ == null) {
        result.session_ = session_;
      } else {
        result.session_ = sessionBuilder_.build();
      }
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }

    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field, java.lang.Object value) {
      return super.setField(field, value);
    }

    @java.lang.Override
    public Builder clearField(com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }

    @java.lang.Override
    public Builder clearOneof(com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }

    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field, int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }

    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field, java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.google.spanner.v1.CreateSessionRequest) {
        return mergeFrom((com.google.spanner.v1.CreateSessionRequest) other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.google.spanner.v1.CreateSessionRequest other) {
      if (other == com.google.spanner.v1.CreateSessionRequest.getDefaultInstance()) return this;
      if (!other.getDatabase().isEmpty()) {
        database_ = other.database_;
        onChanged();
      }
      if (other.hasSession()) {
        mergeSession(other.getSession());
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.google.spanner.v1.CreateSessionRequest parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.google.spanner.v1.CreateSessionRequest) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object database_ = "";
    /**
     *
     *
     * <pre>
     * Required. The database in which the new session is created.
     * </pre>
     *
     * <code>
     * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
     * </code>
     *
     * @return The database.
     */
    public java.lang.String getDatabase() {
      java.lang.Object ref = database_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs = (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        database_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     *
     *
     * <pre>
     * Required. The database in which the new session is created.
     * </pre>
     *
     * <code>
     * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
     * </code>
     *
     * @return The bytes for database.
     */
    public com.google.protobuf.ByteString getDatabaseBytes() {
      java.lang.Object ref = database_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8((java.lang.String) ref);
        database_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     *
     *
     * <pre>
     * Required. The database in which the new session is created.
     * </pre>
     *
     * <code>
     * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
     * </code>
     *
     * @param value The database to set.
     * @return This builder for chaining.
     */
    public Builder setDatabase(java.lang.String value) {
      if (value == null) {
        throw new NullPointerException();
      }

      database_ = value;
      onChanged();
      return this;
    }
    /**
     *
     *
     * <pre>
     * Required. The database in which the new session is created.
     * </pre>
     *
     * <code>
     * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
     * </code>
     *
     * @return This builder for chaining.
     */
    public Builder clearDatabase() {

      database_ = getDefaultInstance().getDatabase();
      onChanged();
      return this;
    }
    /**
     *
     *
     * <pre>
     * Required. The database in which the new session is created.
     * </pre>
     *
     * <code>
     * string database = 1 [(.google.api.field_behavior) = REQUIRED, (.google.api.resource_reference) = { ... }
     * </code>
     *
     * @param value The bytes for database to set.
     * @return This builder for chaining.
     */
    public Builder setDatabaseBytes(com.google.protobuf.ByteString value) {
      if (value == null) {
        throw new NullPointerException();
      }
      checkByteStringIsUtf8(value);

      database_ = value;
      onChanged();
      return this;
    }

    private com.google.spanner.v1.Session session_;
    private com.google.protobuf.SingleFieldBuilderV3<
            com.google.spanner.v1.Session,
            com.google.spanner.v1.Session.Builder,
            com.google.spanner.v1.SessionOrBuilder>
        sessionBuilder_;
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     *
     * @return Whether the session field is set.
     */
    public boolean hasSession() {
      return sessionBuilder_ != null || session_ != null;
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     *
     * @return The session.
     */
    public com.google.spanner.v1.Session getSession() {
      if (sessionBuilder_ == null) {
        return session_ == null ? com.google.spanner.v1.Session.getDefaultInstance() : session_;
      } else {
        return sessionBuilder_.getMessage();
      }
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public Builder setSession(com.google.spanner.v1.Session value) {
      if (sessionBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        session_ = value;
        onChanged();
      } else {
        sessionBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public Builder setSession(com.google.spanner.v1.Session.Builder builderForValue) {
      if (sessionBuilder_ == null) {
        session_ = builderForValue.build();
        onChanged();
      } else {
        sessionBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public Builder mergeSession(com.google.spanner.v1.Session value) {
      if (sessionBuilder_ == null) {
        if (session_ != null) {
          session_ =
              com.google.spanner.v1.Session.newBuilder(session_).mergeFrom(value).buildPartial();
        } else {
          session_ = value;
        }
        onChanged();
      } else {
        sessionBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public Builder clearSession() {
      if (sessionBuilder_ == null) {
        session_ = null;
        onChanged();
      } else {
        session_ = null;
        sessionBuilder_ = null;
      }

      return this;
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public com.google.spanner.v1.Session.Builder getSessionBuilder() {

      onChanged();
      return getSessionFieldBuilder().getBuilder();
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    public com.google.spanner.v1.SessionOrBuilder getSessionOrBuilder() {
      if (sessionBuilder_ != null) {
        return sessionBuilder_.getMessageOrBuilder();
      } else {
        return session_ == null ? com.google.spanner.v1.Session.getDefaultInstance() : session_;
      }
    }
    /**
     *
     *
     * <pre>
     * The session to create.
     * </pre>
     *
     * <code>.google.spanner.v1.Session session = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
            com.google.spanner.v1.Session,
            com.google.spanner.v1.Session.Builder,
            com.google.spanner.v1.SessionOrBuilder>
        getSessionFieldBuilder() {
      if (sessionBuilder_ == null) {
        sessionBuilder_ =
            new com.google.protobuf.SingleFieldBuilderV3<
                com.google.spanner.v1.Session,
                com.google.spanner.v1.Session.Builder,
                com.google.spanner.v1.SessionOrBuilder>(
                getSession(), getParentForChildren(), isClean());
        session_ = null;
      }
      return sessionBuilder_;
    }

    @java.lang.Override
    public final Builder setUnknownFields(final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }

    // @@protoc_insertion_point(builder_scope:google.spanner.v1.CreateSessionRequest)
  }

  // @@protoc_insertion_point(class_scope:google.spanner.v1.CreateSessionRequest)
  private static final com.google.spanner.v1.CreateSessionRequest DEFAULT_INSTANCE;

  static {
    DEFAULT_INSTANCE = new com.google.spanner.v1.CreateSessionRequest();
  }

  public static com.google.spanner.v1.CreateSessionRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CreateSessionRequest> PARSER =
      new com.google.protobuf.AbstractParser<CreateSessionRequest>() {
        @java.lang.Override
        public CreateSessionRequest parsePartialFrom(
            com.google.protobuf.CodedInputStream input,
            com.google.protobuf.ExtensionRegistryLite extensionRegistry)
            throws com.google.protobuf.InvalidProtocolBufferException {
          return new CreateSessionRequest(input, extensionRegistry);
        }
      };

  public static com.google.protobuf.Parser<CreateSessionRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CreateSessionRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.spanner.v1.CreateSessionRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }
}