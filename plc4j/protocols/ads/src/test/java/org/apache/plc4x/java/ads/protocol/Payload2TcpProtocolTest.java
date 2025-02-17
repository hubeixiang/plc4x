/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.ads.protocol;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandlerContext;
import org.apache.plc4x.java.ads.api.tcp.AmsTCPPacket;
import org.apache.plc4x.java.ads.api.tcp.types.UserData;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Collection;
import java.util.stream.Collectors;

import static org.hamcrest.Matchers.hasSize;
import static org.hamcrest.Matchers.instanceOf;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.RETURNS_DEEP_STUBS;
import static org.mockito.Mockito.mock;

@RunWith(Parameterized.class)
public class Payload2TcpProtocolTest extends AbstractProtocolTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(Payload2TcpProtocolTest.class);

    private Payload2TcpProtocol SUT;

    private ChannelHandlerContext channelHandlerContextMock;

    @Parameterized.Parameter
    public AmsTCPPacket amsTCPPacket;

    @Parameterized.Parameter(1)
    public String clazzName;

    @Parameterized.Parameter(2)
    public byte[] amsPacketBytes;

    @Parameterized.Parameters(name = "{index} {1}")
    public static Collection<Object[]> data() {
        return amsPacketStream()
            .map(amsPacket -> new Object[]{
                AmsTCPPacket.of(UserData.of(amsPacket.getByteBuf())),
                amsPacket.getClass().getSimpleName(),
                amsPacket.getBytes()
            })
            .collect(Collectors.toList());
    }

    @Before
    public void setUp() throws Exception {
        SUT = new Payload2TcpProtocol();
        channelHandlerContextMock = mock(ChannelHandlerContext.class, RETURNS_DEEP_STUBS);
        byte[] bytes = amsTCPPacket.getBytes();
        LOGGER.info("amsPacket:\n{} has \n{}bytes\nHexDump:\n{}", amsTCPPacket, bytes.length, amsTCPPacket.dump());
    }

    @Test
    public void encode() throws Exception {
        ArrayList<Object> out = new ArrayList<>();
        SUT.encode(channelHandlerContextMock, Unpooled.wrappedBuffer(amsPacketBytes), out);
        assertEquals(1, out.size());
        assertThat(out, hasSize(1));
    }

    @Test
    public void decode() throws Exception {
        ArrayList<Object> out = new ArrayList<>();
        SUT.decode(channelHandlerContextMock, amsTCPPacket.getByteBuf(), out);
        assertThat(out, hasSize(1));
    }

    @Test
    public void fragmentedDecode() throws Exception {
        ArrayList<Object> out = new ArrayList<>();
        ByteBuf byteBuf = amsTCPPacket.getByteBuf();
        SUT.decode(channelHandlerContextMock, byteBuf.readBytes(byteBuf.readableBytes() - 2), out);
        assertThat(out, hasSize(0));
        SUT.decode(channelHandlerContextMock, byteBuf, out);
        assertThat(out, hasSize(1));
        byteBuf.release();
    }

    @Test
    public void multipleMessagesDecode() throws Exception {
        ArrayList<Object> out = new ArrayList<>();
        ByteBuf byteBuf = Unpooled.wrappedBuffer(amsTCPPacket.getByteBuf(), amsTCPPacket.getByteBuf());
        SUT.decode(channelHandlerContextMock, byteBuf, out);
        assertThat(out, hasSize(2));
        byteBuf.release();
    }

    @Test
    public void roundTrip() throws Exception {
        ArrayList<Object> outbound = new ArrayList<>();
        SUT.encode(channelHandlerContextMock, Unpooled.wrappedBuffer(amsPacketBytes), outbound);
        assertEquals(1, outbound.size());
        assertThat(outbound, hasSize(1));
        assertThat(outbound.get(0), instanceOf(ByteBuf.class));
        ByteBuf byteBuf = (ByteBuf) outbound.get(0);
        ArrayList<Object> inbound = new ArrayList<>();
        SUT.decode(channelHandlerContextMock, byteBuf, inbound);
        assertEquals(1, inbound.size());
        assertThat(inbound, hasSize(1));
        assertThat(inbound.get(0), instanceOf(ByteBuf.class));
        ByteBuf inboundAmsPacket = (ByteBuf) inbound.get(0);
        assertByteBufferEquals(Unpooled.wrappedBuffer(this.amsPacketBytes), inboundAmsPacket);
    }
}